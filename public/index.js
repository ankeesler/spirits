window.onload = () => {
  const inputTextarea = document.getElementById('input');
  const outputTextarea = document.getElementById('output');
  const errorMessage = document.getElementById('error-message');
  const ws = new WebSocket('ws://' + window.location.host + '/api/battle');
  ws.onopen = () => {
    console.log('websocket opened');
  };
  ws.onclose = () => {
    console.log('websocket closed');
  };
  ws.onerror = (e) => {
    console.log(`websocket error: ${e}`);
  };

  const setOutputTextarea = (pending, text) => {
    if (pending) {
      outputTextarea.style = 'color: gray; text-align: center; font-size: 64px';
      outputTextarea.value = '⌛';
    } else {
      outputTextarea.style = 'color: black';
      outputTextarea.value = text;
      outputTextarea.scrollTop = outputTextarea.scrollHeight;
    }
  };

  const setErrorMessage = (message) => {
    errorMessage.innerText = message;
  };

  const runBattle = () => {
    const battle = {type: 'battle-start', details: {spirits: JSON.parse(inputTextarea.value)}};
    ws.send(JSON.stringify(battle));

    ws.onmessage = (m) => {
      console.log(`received websocket message`);
  
      const message = JSON.parse(m.data);
      if (message.type === 'battle-stop') {
        setOutputTextarea(false, message.details.output);
      } else if (message.type === 'action-req') {
        document.getElementById('actions').innerHTML = '';
        outputTextarea.value += message.details.output;

        const spirit = message.details.spirit;
        let actionsHTML = `please select action for spirit ${spirit.name} <br\>`;

        spirit.actions.forEach((a) => actionsHTML += `<button id="action-${a}">${a}</button> <br\>`);
        document.getElementById('actions').innerHTML = actionsHTML;

        spirit.actions.forEach((a) => {
          document.getElementById(`action-${a}`).onclick = (e) => {
            const actionRsp = {type: 'action-rsp', details: {ID: e.target.innerText}};
            ws.send(JSON.stringify(actionRsp));
          };
        });
      } else if (message.type === 'error') {
        setErrorMessage(message.details.reason);
      } else {
        console.log(`received unknown message type: ${message.type}`);
      }
    };
  };

  let timer = setTimeout(() => { }, 0);
  inputTextarea.oninput = (e) => {
    setErrorMessage('');
    setOutputTextarea(true);
    clearTimeout(timer);
    timer = setTimeout(runBattle, 2000);
  };

  document.getElementById('generate-spirits').onclick = (e) => {
    fetch('/api/spirit'+window.location.search, {
      method: 'POST',
    }).then((response) => {
      console.log(`POST ${response.url} response: ${response.status} ${response.statusText}`);

      if (response.status === 200) {
        return response.json()
      } else {
        response.text().then(text => console.log(`POST /api/spirit error: ${text}`));
      }
    }).then((json) => {
      inputTextarea.value = JSON.stringify(json, null, 2) + "\n";
      runBattle();
    }).catch((error) => {
      console.log(`POST /api/spirit error: ${error.message}`);
    });
  };
};
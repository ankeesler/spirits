window.onload = () => {
  const inputTextarea = document.getElementById('input');
  const outputTextarea = document.getElementById('output');
  const errorMessage = document.getElementById('error-message');

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
    fetch('/api/battles', {
      method: 'POST',
      body: inputTextarea.value,
    }).then((response) => {
      console.log(`POST /api/batles response: ${response.status} ${response.statusText}`);

      if (response.status === 200) {
        return response.text()
      } else if (response.status >= 400 && response.status <= 499) {
        response.text().then(text => console.log(`error: ${text}`));
        throw new Error(`invalid spirits json`);
      } else {
        throw new Error('server error :(');
      }
    }).then((text) => {
      setErrorMessage('');
      setOutputTextarea(false, text);
    }).catch((error) => {
      setErrorMessage(error.message);
      setOutputTextarea(false, '');
    });
  };

  let timer = setTimeout(() => { }, 0);
  inputTextarea.oninput = (e) => {
    setErrorMessage('');
    setOutputTextarea(true);
    clearTimeout(timer);
    timer = setTimeout(runBattle, 2000);
  };
};
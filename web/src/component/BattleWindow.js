import React from 'react';
import PropTypes from 'prop-types';

import BattleConsole from './BattleConsole';
import BattleScreen from './BattleScreen';
import log from './../lib/log';

import './BattleWindow.css';

const BattleWindow = (props) => {
  const [output, setOutput] = React.useState('');
  const [timer, setTimer] = React.useState(setTimeout(() => {}, 0));

  const runBattle = (spirits) => {
    log('running battle with ' + spirits);
    props.client.startBattle(spirits, (error, newOutput) => {
      if (error) {
        setOutput(`error: ${error}`);
        return;
      }
      setOutput(newOutput);
    });
  };

  const onSpirits = (spirits) => {
    setOutput('⏳');
    clearTimeout(timer);
    setTimer(setTimeout(() => runBattle(spirits), 2000));
  };

  return (
    <div className="component-battle-window">
      <BattleScreen client={props.client} onSpirits={onSpirits} />
      <BattleConsole message={output} />
    </div>
  );
};

BattleWindow.defaultProps = {
};

BattleWindow.propTypes = {
  client: PropTypes.object.isRequired,
};

export default BattleWindow;

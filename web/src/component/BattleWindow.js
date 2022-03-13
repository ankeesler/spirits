import React from 'react';
import PropTypes from 'prop-types';

import BattleConsole from './BattleConsole';
import BattleScreen from './BattleScreen';
import log from './../lib/log';

import './BattleWindow.css';

const BattleWindow = (props) => {
  const [output, setOutput] = React.useState('');
  const [actioningSpirit, setActioningSpirit] = React.useState(null);

  const handleBattleResolve = (details) => {
    setOutput(details.output);
    if (details.spirit) {
      // This is an action request.
      setActioningSpirit(details.spirit);
    } else {
      // No action request - we are done.
      setActioningSpirit(null);
    }
  };

  const handleBattleReject = (error) => {
    setOutput(`error: ${error}`);
  };

  const startBattle = () => {
    log('running battle with ' + JSON.stringify(props.spirits));
    props.battle
      .start(props.spirits)
      .then(handleBattleResolve)
      .catch(handleBattleReject);
  };

  const stopBattle = () => {
    props.battle
      .stop()
      .then(setActioningSpirit(null))
      .catch(handleBattleReject);
  };

  const onAction = (action) => {
    props.battle
      .action(actioningSpirit, action)
      .then(handleBattleResolve)
      .catch(handleBattleReject);
  };

  return (
    <div className="component-battle-window">
      <button onClick={startBattle}>start</button>
      <button onClick={stopBattle}>stop</button>
      <BattleScreen output={output} />
      <BattleConsole actioningSpirit={actioningSpirit} onAction={onAction} />
    </div>
  );
};

BattleWindow.defaultProps = {
};

BattleWindow.propTypes = {
  battle: PropTypes.object.isRequired,
};

export default BattleWindow;

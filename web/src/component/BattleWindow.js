import React from 'react';
import PropTypes from 'prop-types';

import BattleConsole from './BattleConsole';
import BattleScreen from './BattleScreen';
import log from './../lib/log';

const BattleWindow = (props) => {
  const [inBattle, setInBattle] = React.useState(false);
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
      setInBattle(false);
    }
  };

  const handleBattleReject = (error) => {
    setOutput(`error: ${error}`);
    setInBattle(false);
  };

  const startBattle = () => {
    log('running battle with ' + JSON.stringify(props.spirits));
    setInBattle(true);
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
    <div>
      <div className='row p-2'>
        <div className='col'><button className="btn btn-secondary" id='start-battle-button' onClick={startBattle} disabled={inBattle}>start</button></div>
        <div className='col'><button className="btn btn-secondary" onClick={stopBattle} disabled={!inBattle}>stop</button></div>
      </div>
      <div className='row'>
        <BattleScreen output={output} />
      </div>
      <div className='row'>
        <BattleConsole actioningSpirit={actioningSpirit} onAction={onAction} />
      </div>
    </div>
  );
};

BattleWindow.defaultProps = {
};

BattleWindow.propTypes = {
  battle: PropTypes.object.isRequired,
};

export default BattleWindow;

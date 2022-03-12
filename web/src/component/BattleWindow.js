import React from 'react';
import PropTypes from 'prop-types';

import BattleConsole from './BattleConsole';
import BattleScreen from './BattleScreen';
import log from './../lib/log';

import './BattleWindow.css';

const BattleWindow = (props) => {
  const [output, setOutput] = React.useState('');

  const runBattle = () => {
    log('running battle with ' + JSON.stringify(props.spirits));
    props.client.startBattle(props.spirits, (error, newOutput) => {
      if (error) {
        setOutput(`error: ${error}`);
        return;
      }
      setOutput(newOutput);
    });
  };
  React.useEffect(runBattle);

  return (
    <div className="component-battle-window">
      <BattleScreen output={output} />
      <BattleConsole />
    </div>
  );
};

BattleWindow.defaultProps = {
};

BattleWindow.propTypes = {
  client: PropTypes.object.isRequired,
};

export default BattleWindow;

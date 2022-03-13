import PropTypes from 'prop-types';
import React from 'react';

import Battle from './lib/Battle';
import BattleWindow from './component/BattleWindow';
import Generator from './lib/Generator';
import Header from './component/Header';
import Navigation from './component/Navigation';
import SpiritWindow from './component/SpiritWindow';
import Window from './component/Window';

import './App.css';

function App(props) {
  const [location, setLocation] = React.useState('spirit');
  const [spirits, setSpirits] = React.useState([]);
  return (
    <div className="component-app">
      <Header />
      <Navigation locations={{spirit: true, battle: spirits.length === 2}} onLocation={setLocation} />
      <Window active={location === 'spirit'}>
        <SpiritWindow generator={new Generator(props.client)} onSpirits={setSpirits} />
      </Window>
      <Window active={location === 'battle'}>
        <BattleWindow battle={new Battle(props.client)} spirits={spirits} />
      </Window>
    </div>
  );
};

App.defaultProps = {
};

App.propTypes = {
  client: PropTypes.object.isRequired,
};

export default App;

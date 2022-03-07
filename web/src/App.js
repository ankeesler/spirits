import PropTypes from 'prop-types';

import BattleWindow from './component/BattleWindow';
import Header from './component/Header';

import './App.css';

function App(props) {
  return (
    <div className="component-app">
      <Header />
      <BattleWindow client={props.client} />
    </div>
  );
};

App.defaultProps = {
};

App.propTypes = {
  client: PropTypes.object.isRequired,
};

export default App;

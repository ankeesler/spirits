import PropTypes from 'prop-types';

import SpiritInput from './SpiritInput'

import './BattleScreen.css';

const BattleScreen = (props) => {
  return (
    <div className="component-battle-screen">
      <SpiritInput onSpirits={props.onSpirits} client={props.client} />
    </div>
  );
};

BattleScreen.defaultProps = {
};

BattleScreen.propTypes = {
  client: PropTypes.object.isRequired,
  onSpirits: PropTypes.func.isRequired,
};

export default BattleScreen;

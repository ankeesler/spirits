import PropTypes from 'prop-types';

import SpiritInput from './SpiritInput'

import './BattleScreen.css';

const BattleScreen = (props) => {
  return (
    <div className="component-battle-screen">
      <SpiritInput onSpirits={props.onSpirits} generateSpirits={props.generateSpirits} />
    </div>
  );
};

BattleScreen.defaultProps = {
};

BattleScreen.propTypes = {
  generateSpirits: PropTypes.func.isRequired,
  onSpirits: PropTypes.func.isRequired,
};

export default BattleScreen;

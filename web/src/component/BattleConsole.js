import React from 'react';
import PropTypes from 'prop-types';

import './BattleConsole.css';

const BattleConsole = (props) => {
  const ref = React.useRef(null);
  React.useEffect(() => {
    ref.current.scrollTop = ref.current.scrollHeight;
  });

  const renderMessage = () => {
    if (props.actioningSpirit) {
      return <div style={{padding: '1%'}}>{`select action for spirit ${props.actioningSpirit.name}`}</div>
    }
  };

  const renderButtons = () => {
    if (props.actioningSpirit) {
      return props.actioningSpirit.actions.map((action, i) => {
        return <button key={i} onClick={(e) => props.onAction(e.target.value)}>{action}</button>
      });
    }
  };
  return (
    <div ref={ref} className='component-battle-console'>
      {renderMessage()}
      {renderButtons()}
    </div>
  );
};

BattleConsole.defaultProps = {
};

BattleConsole.propTypes = {
  actioningSpirit: PropTypes.object,
  onAction: PropTypes.func.isRequired,
};

export default BattleConsole;

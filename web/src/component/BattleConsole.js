import React from 'react';
import PropTypes from 'prop-types';

const BattleConsole = (props) => {
  const ref = React.useRef(null);
  React.useEffect(() => {
    ref.current.scrollTop = ref.current.scrollHeight;
  });

  const renderMessage = () => {
    if (props.actioningSpirit) {
      return <div className='padded'>{`select action for spirit ${props.actioningSpirit.name}`}</div>
    }
  };

  const renderButtons = () => {
    if (props.actioningSpirit) {
      return props.actioningSpirit.actions.map((action, i) => {
        return <button className='button' key={i} onClick={(e) => props.onAction(e.target.value)}>{action}</button>
      });
    }
  };
  return (
    <div ref={ref} className='container flex-1 padded border-small'>
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

import React from 'react';
import PropTypes from 'prop-types';

import './BattleConsole.css';

const BattleConsole = (props) => {
  const ref = React.useRef(null);
  React.useEffect(() => {
    ref.current.scrollTop = ref.current.scrollHeight;
  });
  return (
    <div ref={ref} className='component-battle-console'>
      {props.message}  
    </div>
  );
};

BattleConsole.defaultProps = {
  message: '...',
};

BattleConsole.propTypes = {
  message: PropTypes.string,
};

export default BattleConsole;

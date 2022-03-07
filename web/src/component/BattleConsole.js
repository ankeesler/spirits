import React from 'react';
import PropTypes from 'prop-types';

import './BattleConsole.css';

const BattleConsole = (props) => {
  const textareaRef = React.useRef(null);
  React.useEffect(() => {
    textareaRef.current.scrollTop = textareaRef.current.scrollHeight;
  });
  return (
    <textarea ref={textareaRef} className="component-battle-console" defaultValue={props.message} />
  );
};

BattleConsole.defaultProps = {
  message: '...',
};

BattleConsole.propTypes = {
  message: PropTypes.string,
};

export default BattleConsole;

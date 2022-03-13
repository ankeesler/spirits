import React from 'react';
import PropTypes from 'prop-types';

const BattleScreen = (props) => {
  const ref = React.useRef(null);
  React.useEffect(() => {
    ref.current.scrollTop = ref.current.scrollHeight;
  });

  return (
    <div ref={ref} className="p-2">
      <pre id='battle-text'>{props.output}</pre>
    </div>
  );
};

BattleScreen.defaultProps = {
  output: '...',
};

BattleScreen.propTypes = {
  output: PropTypes.string,
};

export default BattleScreen;

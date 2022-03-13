import React from 'react';
import PropTypes from 'prop-types';

const BattleScreen = (props) => {
  const ref = React.useRef(null);
  React.useEffect(() => {
    ref.current.scrollTop = ref.current.scrollHeight;
  });

  return (
    <div ref={ref} className='container container-vertical container-text flex-2 padded' id='battle-text'>
      {props.output}
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

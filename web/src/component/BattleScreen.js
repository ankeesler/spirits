import React from 'react';
import PropTypes from 'prop-types';

import './BattleScreen.css';

const BattleScreen = (props) => {
  const ref = React.useRef(null);
  React.useEffect(() => {
    ref.current.scrollTop = ref.current.scrollHeight;
  });

  return (
    <div ref={ref} className="component-battle-screen">
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

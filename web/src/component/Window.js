import PropTypes from 'prop-types';

import './Window.css';

function Window(props) {
  let className = 'component-window';
  if (!props.active) {
    className += ' hidden';
  }

  return (
    <div className={className}>
      {props.children}
    </div>
  );
};

Window.defaultProps = {
  active: true,
};

Window.propTypes = {
  active: PropTypes.bool.isRequired,
};

export default Window;

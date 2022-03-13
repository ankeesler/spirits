import PropTypes from 'prop-types';

function Window(props) {
  let className = 'container container-vertical border padded';
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

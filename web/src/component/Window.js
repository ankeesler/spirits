import PropTypes from 'prop-types';

function Window(props) {
  let className = 'row bg-light p-1 border border-4';
  if (!props.active) {
    className += ' d-none';
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

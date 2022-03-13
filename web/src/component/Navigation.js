import PropTypes from 'prop-types';

function Navigation(props) {
  const onClick = (e, location) => {
    props.onLocation(location);
  };

  const renderButtons = () => {
    const buttons = [];
    let i = 0;
    for (const location in props.locations) {
      const disabled = !props.locations[location];
      const button = <button key={i++} className='button' id={`show-${location}-window-button`} onClick={(e) => onClick(e, location)} disabled={disabled}>{location}</button>
      buttons.push(button);
    }
    return buttons;
  };
  return (
    <nav>
      {renderButtons()}
    </nav>
  );
};

Navigation.defaultProps = {
};

Navigation.propTypes = {
  locations: PropTypes.object.isRequired,
  onLocation: PropTypes.func.isRequired,
};


export default Navigation;

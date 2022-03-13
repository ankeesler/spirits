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
      const button = <div key={i++} className="col"><button id={location + '-navigation'} className="btn btn-primary" onClick={(e) => onClick(e, location)} disabled={disabled}>{location}</button></div>
      buttons.push(button);
    }
    return buttons;
  };
  return (
    <>
      {renderButtons()}
    </>
  );
};

Navigation.defaultProps = {
};

Navigation.propTypes = {
  locations: PropTypes.object.isRequired,
  onLocation: PropTypes.func.isRequired,
};


export default Navigation;

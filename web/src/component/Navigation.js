import PropTypes from 'prop-types';

import './Navigation.css';

function Navigation(props) {
  const onClick = (e, location) => {
    props.onLocation(location);
  };
  return (
    <div className="component-navigation">
      {props.locations.map((location, i) => (<button key={i} onClick={(e) => onClick(e, location)}>{location}</button>))}
    </div>
  );
};

Navigation.defaultProps = {
};

Navigation.propTypes = {
  locations: PropTypes.array.isRequired,
  onLocation: PropTypes.func.isRequired,
};


export default Navigation;

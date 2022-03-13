import React from 'react';
import PropTypes from 'prop-types';

const Spirit = (props) => {
  return (
    <div className='container container-vertical container-centered'>
      <h4 className='padded'>{props.name}</h4>
      <div className='padded'>{props.health ? `${props.health}/${props.health}` : '?'}</div>
    </div>
  );
};

Spirit.defaultProps = {
};

Spirit.propTypes = {
  name: PropTypes.string.isRequired,
  health: PropTypes.number.isRequired,
};

const Actions = (props) => {
  return (
    <div className='container container-vertical container-centered'>
      <h4>actions</h4>
      {props.ids.map((id, i) => <div key={i}>{id}</div>)}
    </div>
  );
};

Actions.defaultProps = {
};

Actions.propTypes = {
  ids: PropTypes.array.isRequired,
};

const Stat = (props) => {
  return (
    <div className='container container-vertical container-centered'>
      <h4 className='padded'>{props.name}</h4>
      <div className='padded'>{props.value ? props.value : 0}</div>
    </div>
  );
};

Stat.defaultProps = {
};

Stat.propTypes = {
  name: PropTypes.string.isRequired,
  value: PropTypes.number,
};

const SpiritInfo = (props) => {
  return (
    <div className='container border-small margin-bottom'>
      <Spirit name={props.spirit.name} health={props.spirit.health} />
      <Actions ids={props.spirit.actions} />
      <Stat name='power' value={props.spirit.power} />
      <Stat name='armor' value={props.spirit.armor} />
      <Stat name='agility' value={props.spirit.agility} />
    </div>
  );
};

SpiritInfo.defaultProps = {
};

SpiritInfo.propTypes = {
  spirit: PropTypes.object.isRequired,
};

export default SpiritInfo;

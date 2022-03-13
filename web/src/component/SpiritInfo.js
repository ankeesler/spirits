import React from 'react';
import PropTypes from 'prop-types';

import getRandomColor from './../lib/getRandomColor';
import log from './../lib/log';

const RandomImage = (props) => {
  const gradientFrom = getRandomColor(props.seed + ' 0');
  const gradientTo = getRandomColor(props.seed + ' 1');
  const style = {
    background: `linear-gradient(${gradientFrom}, ${gradientTo})`,
    width: '100%',
    height: '100%',
  };
  return (
    <div className='container padded'>
      <div style={style}></div>
    </div>
  );
};

RandomImage.defaultProps = {
};

RandomImage.propTypes = {
  seed: PropTypes.string.isRequired,
};

const Spirit = (props) => {
  const [human, setHuman] = React.useState(false);

  const onChange = (e) => {
    props.onHumanChecked(e.target.checked);
    setHuman(e.target.checked);
  };

  React.useEffect(() => {
    setHuman(props.human);
  }, [props.human]);

  return (
    <div className='container container-vertical container-centered'>
      <RandomImage seed={props.name} />
      <h4>{props.name}</h4>
      <div>{props.health ? `${props.health}/${props.health}` : '?'}</div>
      <div>
        <input type='checkbox' id={`${props.name}-human`} checked={human} onChange={onChange} />
        <label htmlFor={`${props.name}-human`}> human</label>
      </div>
    </div>
  );
};

Spirit.defaultProps = {
};

Spirit.propTypes = {
  name: PropTypes.string.isRequired,
  health: PropTypes.number.isRequired,
  human: PropTypes.bool.isRequired,
  onHumanChecked: PropTypes.func.isRequired,
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
  ids: ['attack'],
};

Actions.propTypes = {
  ids: PropTypes.array,
};

const Stat = (props) => {
  return (
    <div className='container container-vertical container-centered'>
      <h4>{props.name}</h4>
      <div>{props.value ? props.value : 0}</div>
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
  const onHumanChecked = (checked) => {
    props.spirit.intelligence = (checked ? 'human' : 'roundrobin');

    log('set spirit intelligence to ' + props.spirit.intelligence);
  };
  return (
    <div className='container border-small'>
      <Spirit name={props.spirit.name} health={props.spirit.health} human={props.spirit.intelligence === 'human'} onHumanChecked={onHumanChecked} />
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

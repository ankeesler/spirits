import React from 'react';
import PropTypes from 'prop-types';

import log from './../lib/log';

import './SpiritWindow.css';
import './Window.css';

const SpiritWindow = (props) => {
  const [spirits, setSpirits] = React.useState('');

  const onSpirits = (spirits) => {
    setSpirits(spirits);

    try {
      props.onSpirits(JSON.parse(spirits));
    } catch (error) {
      log(`error parsing spirits: ${error}`);
      props.onSpirits([]);
    }
  };

  const onClick = async (e) => {
    props.generator.generate().then((generatedSpirits) => {
      onSpirits(JSON.stringify(generatedSpirits));
    }).catch((error) => {
      log(`generate spirits error: ${error}`);
    });
  };

  return (
    <div className='component-spirit-window'>
      <button onClick={onClick}>generate</button>
      <div onInput={e => onSpirits(e.target.value)}>{spirits}</div>
    </div>    
  );
};

SpiritWindow.defaultProps = {
};

SpiritWindow.propTypes = {
  onSpirits: PropTypes.func.isRequired,
  generator: PropTypes.object.isRequired,
};

export default SpiritWindow;

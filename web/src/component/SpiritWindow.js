import React from 'react';
import PropTypes from 'prop-types';

import log from './../lib/log';

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
      onSpirits(JSON.stringify(generatedSpirits, null, 2));
    }).catch((error) => {
      log(`generate spirits error: ${error}`);
    });
  };

  return (
    <div className='container'>
      <div>
        <button className='button' id='generate-spirits-button' onClick={onClick}>generate</button>
      </div>
      <textarea className='container border padded' id='spirits-text' onInput={e => onSpirits(e.target.value)} value={spirits}></textarea>
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

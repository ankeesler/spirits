import React from 'react';
import PropTypes from 'prop-types';

import log from './../lib/log';
import SpiritInfo from './SpiritInfo.js';

const SpiritWindow = (props) => {
  const [spirits, setSpirits] = React.useState([]);

  const onClick = async (e) => {
    props.generator.generate().then((generatedSpirits) => {
      props.onSpirits(generatedSpirits);
      setSpirits(generatedSpirits);
    }).catch((error) => {
      log(`generate spirits error: ${error}`);
    });
  };

  return (
    <div className='container container-vertical' id='spirit-window'>
      <div>
        <button className='button' onClick={onClick}>generate</button>
      </div>
      {spirits.map((spirit, i) => <SpiritInfo key={i} spirit={spirit} />)}
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

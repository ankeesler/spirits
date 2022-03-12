import React from 'react';
import PropTypes from 'prop-types';

import log from './../lib/log';

import './SpiritInput.css';

const SpiritInput = (props) => {
  const [spirits, setSpirits] = React.useState('');

  const onSpirits = (e) => {
    setSpirits(e.target.value);
    try {
      props.onSpirits(JSON.parse(e.target.value));
    } catch (error) {
      log(`invalid spirits: ${error}`);
      props.onSpirits([]);
    }
  };

  const onDoubleClick = async (e) => {
    props.client.generateSpirits((error, generatedSpirits) => {
      if (error) {
        log(`generate spirits error: ${error}`);
        return;
      }
      setSpirits(JSON.stringify(generatedSpirits));
      props.onSpirits(generatedSpirits);
    });
  };

  return (
    <textarea className="component-spirit-input" onInput={onSpirits} onDoubleClick={onDoubleClick} value={spirits} placeholder="enter spirits JSON; double click to generate" />
  );
};

SpiritInput.defaultProps = {
};

SpiritInput.propTypes = {
  onSpirits: PropTypes.func.isRequired,
  client: PropTypes.object.isRequired,
};

export default SpiritInput;

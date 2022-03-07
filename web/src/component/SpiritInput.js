import React from 'react';
import PropTypes from 'prop-types';

import './SpiritInput.css';

const SpiritInput = (props) => {
  const [spirits, setSpirits] = React.useState('');

  const onSpirits = (e) => {
    setSpirits(e.target.value);
    props.onSpirits(e.target.value);
  };

  const onDoubleClick = async (e) => {
    props.generateSpirits((error, generatedSpirits) => {
      setSpirits(generatedSpirits);
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
  generateSpirits: PropTypes.func.isRequired,
};

export default SpiritInput;

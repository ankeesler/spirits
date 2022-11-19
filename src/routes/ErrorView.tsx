import React, {FC} from 'react';
import Alert from 'react-bootstrap/Alert';
import {Container} from 'react-bootstrap';
import {useRouteError} from 'react-router-dom';

interface ErrorViewProps {

};

// interface Error {
//   statusText: string
//   message: string
// }

const ErrorView: FC<ErrorViewProps> = (props) => {
  const error = useRouteError() as Error;
  return (
    <Container className="mt-3">
      <Alert variant='danger'>
        <p>{error.message}</p>
        <p><pre>{error.stack}</pre></p>
      </Alert>
    </Container>
  );
};

export default ErrorView;

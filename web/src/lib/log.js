const log = (s) => {
  if (process.env.NODE_ENV === 'development') {
    console.log(s);
  }
};

export default log;
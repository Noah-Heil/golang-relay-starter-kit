import 'babel/polyfill';

import App from './components/App';
import AppHomeRoute from './routes/AppHomeRoute';
import React from 'react';
import ReactDOM from 'react-dom';
import Relay from 'react-relay';

// import Button from './components/button';
import Kitten from './image';

// console.log(Button);

var newMessage = () => (`
  <p>
    ${Kitten}
  </p>
`);


ReactDOM.render(
  <Relay.RootContainer Component={App} route={new AppHomeRoute()}/>,
  document.getElementById('root')
);

// var app = document.getElementById('app');
// app.innerHTML = newMessage();

// Button.attachEl();

if(module.hot) {
  module.hot.accept();
}
import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import './index.css';

ReactDOM.render(
  <App />,
  document.getElementById('root')
);

document.getElementById("baca-submit").addEventListener("click", function(e){
  e.preventDefault();
  document.getElementById("result").removeAttribute("hidden");
})

document.getElementById("eja-submit").addEventListener("click", function(e){
  e.preventDefault();
  document.getElementById("result").removeAttribute("hidden");
})
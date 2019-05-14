import React, { Component } from 'react';
import './App.css';
import axios from 'axios'

class App extends Component {
  constructor(){
    super()
    this.ejaClick = this.ejaClick.bind(this)
    this.bacaClick = this.bacaClick.bind(this)
  }

  ejaClick(){
    //GET
    axios.get('http://0.0.0.0:8080/spell?number='+document.getElementById("eja").value,{  crossDomain: true
    }).then(response => document.getElementById("result").innerHTML = 
      "<span>Hasil Pengejaan</span> = "+response.data.text)
  }

  bacaClick(){
    //POST
    axios.post('http://0.0.0.0:8080/read?text='+document.getElementById("baca").value,{
      crossDomain: true
    }).then(response => document.getElementById("result").innerHTML = 
    "<span>Hasil Pembacaan</span> = "+response.data.number)
  }
  
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <h2>Welcome to Indonesian Numeral Speller</h2>
        </div>
        <div className="App-container">
          <div className="eja">
            <h4>Pengejaan (angka ke teks)</h4>
            <p> Masukkan angka yang valid contohnya 23, 245, 4587. Angka maksimum yang dapat diteima adalah 2 Miliyar</p>
            <form>
              <input type="number" name="eja" placeholder="Masukkan angka" id="eja"></input>
              <button id="eja-submit" onClick={this.ejaClick}>submit</button>
            </form>
          </div>
          <div className="baca">
            <h4>Pembacaan (teks ke angka)</h4>
            <p> Masukkan teks yang valid. Pastikan tidak terjadi typo. Angka maksimum yang dapat diteima adalah 2 Miliyar</p>
            <form>
              <input type="text" name="baca" placeholder="Masukkan teks" id="baca"></input>
              <button id="baca-submit" onClick={this.bacaClick}>submit</button>
            </form>
          </div>
          <div className="result" id="result" hidden>
          </div>
        </div>
        <div className="App-footer">
          <p>Created by : Juniardi Akbar</p>
        </div>
      </div>
    );
  }
}

export default App;

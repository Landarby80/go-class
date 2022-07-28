import axios from "axios";
import React from "react";

import 'bootstrap/dist/css/bootstrap.min.css'
 

// const baseURL = "http://localhost:8000/createUsername";


export default class createUsername extends React.Component {
  state = {
    name: ''
  }

  handleChange = event => {
    this.setState({ name: event.target.value });
    console.log(event.target.value)
  }

  handleSubmit = event => {
    event.preventDefault();

    const user = {
      name: this.state.name
    };

    const json = JSON.stringify( user );//converts a JavaScript object or value to a JSON string
    console.log(user.name)

    axios.post(`http://localhost:8000/createUsername`, json)
      .then(res => {
        console.log(res);
        console.log(res.data);
       
      })
  }

  render() {
    return (
      <div>
        <form onSubmit={this.handleSubmit}>
          <label>
            Username:
            <input type="text" name="name" onChange={this.handleChange} />
          </label>
          <button type="submit">Add</button>
        </form>
      </div>
    )
  }
}

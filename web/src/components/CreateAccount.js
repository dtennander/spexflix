import React, { Component } from 'react';
import Button from "./Button";
import InputField from "./InputField";

const divStyle = {
    textAlign: 'center'
};

const headerStyle = {
    height: '50px',
    padding: '10px',
    color: '#71141b'
};

const titleStyle = {
    fontSize: '1.5em'
};


class CreateAccount extends Component {
    render() {
        return (
        <div style={divStyle}>
            <header style={headerStyle}>
                <h1 style={titleStyle}>Skapa din användare:</h1>
            </header>
            <form>
               <table align="center">
                   <tbody>
                   <tr>
                       <td width="70px"><InputField type="text" name="first" placeholder="Förnamn"/></td>
                       <td width="10"/>
                       <td colSpan="2"><InputField type="text" name="last" placeholder="Efternamn"/></td>
                   </tr>
                   <tr>
                       <td colSpan="4"><InputField type="text" name="email" placeholder="e-post"/></td>
                   </tr>
                   <tr>
                       <td colSpan="4"><InputField type="password" name="password" placeholder="Nytt Lösenord"/></td>
                   </tr>
                   <tr>
                       <td colSpan="3"/><td width="70px"><Button text="skicka"/></td>
                   </tr>
                   </tbody>
               </table>
            </form>
        </div>
        );
    }
}

export default CreateAccount;

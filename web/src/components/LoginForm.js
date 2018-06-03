import React, {Component} from 'react';
import InputField from "./InputField";
import Button from "./Button";
import Api from '../api'

/**
 * @param onSuccessfullLogIn
 */
class LoginForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            loggingIn: false,
        };
        this.email = React.createRef();
        this.password = React.createRef();
        this.login = this.login.bind(this);
    }

    login() {
        const email = this.email.current.getInput();
        const password = this.password.current.getInput();
        this.setState({loggingIn: true});
        Api.LogInUser(email, password)
            .then(this.props.onSuccessfulLogIn)
            .catch(error => {
                console.log(error);
                this.setState({loggingIn: false});
            })
    }

    render() {
        let buttonText;
        if (this.state.loggingIn) {
            buttonText = "Loggar in...";
        } else {
            buttonText = "Logga in";
        }
        return (
            <form method="">
                <table align="center">
                    <tbody>
                    <tr>
                        <td width="100px">
                            <InputField
                                ref={this.email}
                                type="text"
                                name="email"
                                placeholder="e-post"
                                onEnter={this.login}
                            />
                        </td>
                        <td width="100px">
                            <InputField
                                ref={this.password}
                                type="password"
                                name="password"
                                placeholder="Lösenord"
                                onEnter={this.login}
                            />
                        </td>
                        <td>
                            <Button text={buttonText} onClick={this.login}/>
                        </td>
                    </tr>
                    </tbody>
                </table>
            </form>
        )
    }
}

export default LoginForm
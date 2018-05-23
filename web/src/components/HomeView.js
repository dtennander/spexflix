import React, {Component} from 'react';

class HomeView extends Component{
    render() {
        return (
            <h1>Välkommen {this.props.user.name}!</h1>
        )
    }
}

export default HomeView
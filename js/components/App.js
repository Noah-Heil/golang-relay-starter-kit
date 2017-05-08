import React from 'react';
import Relay from 'react-relay';

// React Component
class App extends React.Component {
    render() {
        return ( <div><h1> { this.props.user.firstname } </h1> </div>
        );
    }
}

// Relay Container
export default Relay.createContainer(App, {
    fragments: {
        user: () => Relay.QL`
        fragment on User {
            firstname
        }
        `,
    },
});
import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import CharactersList from './CharactersList'

const App = () => {
    return (
        <div>
            <React.Fragment>
                <CssBaseline />
                <Container maxWidth="md">
                    <CharactersList />
                </Container>
            </React.Fragment>
        </div>
    );
};


export default App;
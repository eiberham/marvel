import React, {useState, useEffect} from 'react';

import marvel from "../apis/marvel";

const CharactersList = () => {
    const [characters, setCharacters] = useState({characters: []});

    useEffect( () => {
        const fetchCharacters = async () => {
            const {data:{data}} = await marvel.get('/characters');
            setCharacters(data);
        }

        fetchCharacters()

    }, [])

    return (
        <div>
            CharactersList Component !
        </div>
    );
};

export default CharactersList;
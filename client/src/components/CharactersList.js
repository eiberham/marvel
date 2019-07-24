import React, {useState, useEffect} from 'react';

import CharacterCard from './CharacterCard';

import marvel from "../apis/marvel";

const CharactersList = () => {
    const [characters, setCharacters] = useState([]);

    useEffect( () => {
        const fetchCharacters = async () => {
            try {
                const {data:{data:{results}}} = await marvel.get('/characters');
                console.log("resultados: ", results);
                setCharacters(results);
            }catch(error) {
                console.error(error);
            }
        }

        fetchCharacters()

    }, []);

    return (
        <section className="characters">
            {characters && characters.map(({name, description, thumbnail:{path, extension}}) => (
                <CharacterCard
                    thumbnail={`${path}.${extension}`}
                    title={name}
                    description={description}
                />
            ))}
        </section>
    );
};

export default CharactersList;
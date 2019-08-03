import React, {useState, useEffect} from 'react';
import CharacterCard from './CharacterCard';
import marvel from "../apis/marvel";
import styled from "@emotion/styled";

const Section = styled.section`
    display: grid;
    grid-template-rows: 200px;
    grid-template-columns: repeat(3, 1fr);
    grid-gap: 30px;
    overflow: hidden;
`;

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
        <Section>
            {characters && characters.map(({name, description, thumbnail:{path, extension}}) => (
                <CharacterCard
                    thumbnail={`${path}.${extension}`}
                    title={name}
                    description={description}
                />
            ))}
        </Section>
    );
};

export default CharactersList;
import React from 'react';
import styled from '@emotion/styled';

const Div = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
`;

const CharacterCard = (props) => {
    return (
        <Div>
            <figure>
                <img src={props.thumbnail} alt={props.name} width={200}/>
                <figcaption>
                    {props.name}
                </figcaption>
            </figure>
        </Div>
    )
};


export default CharacterCard;
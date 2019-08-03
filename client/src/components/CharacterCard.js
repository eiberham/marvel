import React, { useState, useRef } from 'react';
import styled from '@emotion/styled';

const Div = styled.div`
    
`;

const CharacterCard = (props) => {
    const image = useRef(null);
    const [spans, setSpans] = useState(0);

    function getSpans(){
        const height = image.current.clientHeight;
        const spans = Math.ceil(height / 10);
        setSpans(spans);
    };

    return (
        <Div style={{gridRowEnd: `span ${spans}`}}>
            <figure>
                <img src={props.thumbnail} alt={props.name} width={250} ref={image} onLoad={getSpans}/>
                <figcaption>
                    {props.name}
                </figcaption>
            </figure>
        </Div>
    )
};

export default CharacterCard;
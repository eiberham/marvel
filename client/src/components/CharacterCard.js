import React, { useState, useRef } from 'react';
import styled from '@emotion/styled';
import PropTypes from "prop-types";

const Div = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    position: relative;
    margin: 0.5rem;
    figure {
        img {
        height: 90%;
        &:hover {
            cursor: pointer;
            filter: drop-shadow(2px 4px 6px black);
            outline: 0.5rem solid #000;
        }
    }
    
    & > figcaption {
        position: absolute;
        transform: translate(-50%, -50%);
        top: 69%;
        left: 50%;
        font-family: 'Marvel', sans-serif;
        font-weight: 700;
        font-size: 1rem;
        -webkit-text-stroke: 1px black;
        color: white;
        text-shadow:
           3px 3px 0 #000,
          -1px -1px 0 #000,  
           1px -1px 0 #000,
          -1px 1px 0 #000,
           1px 1px 0 #000;
    }
`;

const CharacterCard = (props) => {
    const image = useRef(null);
    const [spans, setSpans] = useState(0);

    const getSpans = () => {
        const height = image.current.clientHeight;
        const spans = Math.ceil(height / 10);
        setSpans(spans);
    };

    return (
        <Div style={{gridRowEnd: `span ${spans}`}}>
            <figure>
                <img
                    src={props.thumbnail}
                    alt={props.title}
                    width={250}
                    ref={image}
                    onLoad={getSpans}
                />
                <figcaption>
                    {props.title}
                </figcaption>
            </figure>
        </Div>
    )
};

CharacterCard.propTypes = {
    thumbnail: PropTypes.string.isRequired,
    title: PropTypes.string
}

export default CharacterCard;
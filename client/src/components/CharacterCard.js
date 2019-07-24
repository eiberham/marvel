import React from 'react';

import styled from '@emotion/styled';

const Section = styled.section`
  color: turquoise;
`;

const CharacterCard = (props) => {
    return (
        <div>
            <figure>
                <figcaption>
                    {props.name}
                </figcaption>
            </figure>
        </div>
    )
};


export default CharacterCard;
import React from "react";
import styled from "@emotion/styled";

const Section = styled.section`
    display: flex;
    justify-content: center;
    align-self: flex-end;
    font-family: 'Marvel', sans-serif;
    font-weight: 700;
    font-size: 2rem;
    background: gray;
`;

function Footer(){
    console.log("render footer component");
    return (
        <Section>
            <footer>Data provided by Marvel. © 2019 MARVEL"</footer>
        </Section>
    )
}

export default Footer;
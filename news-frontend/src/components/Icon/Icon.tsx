import React from "react";
import { Image } from "react-bootstrap";

interface IIConProps {
    src: string;
}

const Icon = ({src}: IIConProps) => <Image src={src} rounded width={32} height={32} />;

export default Icon;
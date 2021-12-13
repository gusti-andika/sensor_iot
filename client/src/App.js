import React from "react";
import Temperature from "./Temperature";

export default () => {
    

    return (<div className="container">
        <Temperature kind="latest" useColor={true} />
        <Temperature kind="min"  />
        <Temperature kind="max"  />
        </div>)
}
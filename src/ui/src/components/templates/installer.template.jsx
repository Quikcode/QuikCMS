import {Outlet} from "react-router-dom";
import s from "./styles/installer.module.scss"
import {Select, MenuItem} from "@mui/material";

export default () => {

    return (


        <>
            <Outlet/>
            <footer className={s.Footer}>
                <p>v1.0.0 Alpha</p>
                <Select
                    value={10}
                >
                    <MenuItem value={10}>Ten</MenuItem>
                    <MenuItem value={20}>Twenty</MenuItem>
                    <MenuItem value={30}>Thirty</MenuItem>
                </Select>
            </footer>
        </>
    )
}
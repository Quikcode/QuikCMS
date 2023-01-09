import React from "react";
import {BrowserRouter, Route, Routes} from "react-router-dom";

import DefaultTemplate from "@/components/templates/default.template";
import InstallerTemplate from "@/components/templates/installer.template";
import InstallerHomePage from "@/pages/installer/home";
import AuthTemplate from "@/components/templates/auth.template";

export default () => {

    return (
        <BrowserRouter>
            <Routes>
                {/* Installer Rotues */}
                <Route path={"/installer"} element={<InstallerTemplate/>}>
                    <Route index element={<InstallerHomePage/>}/>
                </Route>
                
                {/* Auth Routes */}
                <Route path={"/auth"} element={<AuthTemplate/>}>
                    <Route index element={<h1>Auth</h1>}/>
                </Route>

                {/* Global Routes */}
                <Route path={"/"} element={<DefaultTemplate/>}>
                    <Route index element={<h1>Home</h1>}/>
                </Route>
            </Routes>
        </BrowserRouter>
    )
}
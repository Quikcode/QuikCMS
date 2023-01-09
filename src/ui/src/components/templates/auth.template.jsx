import {Outlet} from "react-router-dom";
import {useTranslation} from "react-i18next";

export default () => {
    const [t] = useTranslation("global")

    return (
        <>
            {t("prueba")}
            <Outlet/>
        </>
    )
}
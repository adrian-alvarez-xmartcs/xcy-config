import { useEffect, useRef } from "react";
import AppBundle from '../../consts/appBundle';

const useTitle = (title: string) => {
    const documentDefined = typeof document !== 'undefined';
    const originalTitle = useRef<string | null>(documentDefined ? document.title : null);

    const titleBuilt = `${AppBundle.Name} - ${title}`;

    useEffect(() => {
        if (!documentDefined) return;

        if (document.title !== titleBuilt) document.title = titleBuilt;

        return () => {
            if (originalTitle.current !== null) {
                document.title = originalTitle.current;
            }
        };
    }, [titleBuilt]);
};

export default useTitle;

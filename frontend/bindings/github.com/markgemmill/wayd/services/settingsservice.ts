// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as slog$0 from "../../../../log/slog/models.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./models.js";

export function DatabasePath(): Promise<string> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3851639849) as any;
    return $resultPromise;
}

export function GetSettings(): Promise<$models.Settings | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3291556084) as any;
    let $typingPromise = $resultPromise.then(($result) => {
        return $$createType1($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Initialize(logger: slog$0.Logger | null): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1405771109, logger) as any;
    return $resultPromise;
}

export function Load(logger: slog$0.Logger | null): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(588316455, logger) as any;
    return $resultPromise;
}

export function SetSettings(settings: $models.Settings | null): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3264572568, settings) as any;
    return $resultPromise;
}

export function Write(): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(840163986) as any;
    return $resultPromise;
}

// Private type creation functions
const $$createType0 = $models.Settings.createFrom;
const $$createType1 = $Create.Nullable($$createType0);

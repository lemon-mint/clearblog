window.addEventListener(
    "onload",
    (e) => {
        mdc.drawer.MDCDrawer.attachTo(document.querySelector('.mdc-drawer'));
        mdc.ripple.MDCRipple.attachTo(document.querySelector('.foo-button'));
        const list = mdc.list.MDCList.attachTo(document.querySelector('.mdc-list'));
        list.wrapFocus = true;
    }
);


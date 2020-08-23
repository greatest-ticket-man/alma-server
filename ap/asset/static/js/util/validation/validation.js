'use strict';

class ValidationUtil {
    static email(email) {
        const re = /\S+@\S+\.\S+/;
        return re.test(email);
    }
}



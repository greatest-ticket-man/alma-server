'use strict';

class ValidationUtil {
    static email(email) {
        console.log("email validation ");
        let re = /\S+@\S+\.\S+/;
        return re.test(email);
    }
}



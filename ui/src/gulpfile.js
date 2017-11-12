const gulp = require('gulp'),
      gutil = require('gulp-util'),
      connect = require("gulp-connect"),
      path = require('path');

const rootPath = path.join(__dirname, "src");

const paths = {
    static: rootPath + "/**/*.{html,css,js}" 
}

gulp.task('static', () => {
    gulp.src(paths.static)
        .pipe(connect.reload());
});

gulp.task('watch', () => {
    gulp.watch(paths.static, {
        interval: 1000,
        usePolling: true
    }, ['static']);
});

gulp.task('connect', () => {
    connect.server({
        root: rootPath,
        livereload: {
            hostname: "localhost"
        },
        port: 5050,
        host: "0.0.0.0"
    });
});

gulp.task('default', ['connect', 'watch']);

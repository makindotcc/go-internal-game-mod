/**
 * Expose dyld.h apis to our golang code.
 */

#include <mach-o/dyld.h>

uint64_t GetMainModuleAddress() {
    return (uint64_t)_dyld_get_image_header(0);
}
